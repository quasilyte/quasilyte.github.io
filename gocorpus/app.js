var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var pako;
(function (pako) {
})(pako || (pako = {}));
var App;
(function (App) {
    class RepoData {
        constructor() {
            this.files = [];
        }
    }
    class RepoFileData {
    }
    const appState = {
        metadata: (null),
        corpus: new Map(),
        go: null,
        wasm: null,
        busy: false,
        running: false,
        // Current query state.
        runError: '',
        runStartTime: 0,
        searchResults: new Map(),
        filesTotal: 0,
        filesScanned: 0,
        slocProcessed: 0,
        hits: 0,
    };
    function updateStatus(status) {
        document.getElementById('status').innerText = status;
    }
    function ready() {
        if (appState.runError === '') {
            updateStatus('ready');
        }
        else {
            updateStatus('ERROR ' + appState.runError);
        }
        appState.runError = '';
    }
    function loadMetadata() {
        updateStatus('loading metadata');
        console.log("loadMetadata()");
        return fetch('corpus-output/corpus.json')
            .then(response => response.json());
    }
    function loadWASM(metadata) {
        appState.metadata = metadata;
        let $totalSLOC = document.getElementById('total-sloc');
        let totalSLOC = 0;
        for (let repo of appState.metadata.Repositories) {
            totalSLOC += repo.SLOC;
        }
        $totalSLOC.innerText = totalSLOC.toLocaleString();
        updateStatus('loading wasm');
        console.log("loadWASM()");
        appState.go = new Go();
        return new Promise((resolve, reject) => {
            WebAssembly.instantiateStreaming(fetch("main.wasm"), appState.go.importObject)
                .then(result => resolve(result.instance))
                .catch(error => reject(error));
        });
    }
    function doChunks(count, f, finished) {
        var i = 0;
        (function step() {
            if (!f(i)) {
                finished(true);
                return;
            }
            i++;
            if (i < count) {
                setTimeout(step, 0);
            }
            else {
                finished(false);
            }
        })();
    }
    function getFrequencyStats() {
        // append($_, $_) => 24.1662
        const baselineFrequency = 70.0; // `err != nil` score
        const resultFrequency = appState.slocProcessed / appState.hits;
        const relativeRate = 100.0 * (baselineFrequency / resultFrequency);
        return [relativeRate, Math.round(resultFrequency)];
    }
    function htmlEscape(s) {
        let div = document.createElement('div');
        div.innerText = s;
        return div.innerHTML;
    }
    function searchDone() {
        let endTime = window.performance.now();
        let elapsedMillis = endTime - appState.runStartTime;
        let elapsedSeconds = elapsedMillis / 1000.0;
        appState.busy = false;
        ready();
        var $run = document.getElementById('run-button');
        $run.innerText = 'Run';
        appState.running = false;
        let $progress = document.getElementById('search-progress');
        $progress.innerHTML = `Progress: 100% (hits: ${appState.hits})`;
        var $results = document.getElementById('search-results');
        var parts = [];
        var sortedMatches = [...appState.searchResults.entries()].sort((a, b) => b[1] - a[1]);
        if (appState.hits !== 0) {
            let [freqScore, freqPerSLOC] = getFrequencyStats();
            parts.push(`<p><i>Frequency score: ${freqScore.toFixed(4)} (1 match per ~${freqPerSLOC.toLocaleString()} SLOC)</i></p>`);
        }
        else {
            parts.push(`<p><i>Frequency score: 0</i></p>`);
        }
        parts.push(`<p><i>Time elapsed: ${elapsedSeconds.toFixed(2)} sec</i></p>`);
        for (let e of sortedMatches) {
            let [m, num] = e;
            let numStr = num == 1 ? '' : ` (${num} matches)`;
            parts.push(`<li><span class="result">${htmlEscape(m)}${numStr}</span></li>`);
        }
        $results.innerHTML += '<ol>' + parts.join('') + '</ol>';
    }
    function runQueryRecursive(pattern, filter, toScan) {
        if (toScan.length == 0) {
            searchDone();
            return;
        }
        let repo = toScan.pop();
        let repoData = appState.corpus.get(repo.Name);
        let files = repoData.files;
        doChunks(files.length, i => {
            if (!appState.running) {
                return false;
            }
            let f = files[i];
            let filename = f.name;
            const maxLength = 50;
            if (filename.length > maxLength) {
                let nameParts = filename.split('/');
                let baseName = nameParts[nameParts.length - 1];
                let shorterName = '';
                if (baseName.length >= maxLength) {
                    shorterName = '{...}/' + baseName;
                }
                else {
                    let prefix = nameParts.slice(0, -1).join('/');
                    let prefixLen = maxLength - baseName.length - 4;
                    shorterName = prefix.substr(0, prefixLen) + '{...}/' + baseName;
                }
                filename = shorterName;
            }
            updateStatus(`processing ${filename}`);
            let $progress = document.getElementById('search-progress');
            let progressValue = Math.round((appState.filesScanned / appState.filesTotal) * 100);
            $progress.innerHTML = `Progress: ${progressValue}% (hits: ${appState.hits})`;
            let fileInfo = repo.Files[i];
            let result = gogrep({
                pattern: pattern,
                filter: filter,
                fileFlags: fileInfo.Flags,
                fileMaxDepth: fileInfo.MaxDepth,
                targetName: f.name,
                targetSrc: f.contents,
            });
            if (result.err) {
                console.error(`grepping ${f.name}: ${result.err}`);
                let canContinue = result.err.includes('parse Go:');
                if (canContinue) {
                    return true;
                }
                appState.runError = result.err;
                appState.running = false;
                return false;
            }
            if (!result.skipped) {
                appState.slocProcessed += fileInfo.SLOC;
            }
            appState.filesScanned++;
            if (!result.matches) {
                return true;
            }
            appState.hits += result.matches.length;
            for (let m of result.matches) {
                if (appState.searchResults.size >= 1000) {
                    break;
                }
                if (appState.searchResults.has(m)) {
                    appState.searchResults.set(m, appState.searchResults.get(m) + 1);
                }
                else {
                    appState.searchResults.set(m, 1);
                }
            }
            return true;
        }, (stopped) => {
            if (stopped) {
                searchDone();
            }
            else {
                runQueryRecursive(pattern, filter, toScan);
            }
        });
    }
    function loadRepo(repo) {
        updateStatus(`loading ${repo.Name} repository...`);
        return fetch(`corpus-output/${repo.Name}.tar.gz`).
            then(result => new Promise((resolve, reject) => {
            result.arrayBuffer().
                then(b => resolve({ repo: repo, archive: new Uint8Array(b) }));
        }));
    }
    function loadRecursive(toLoad, onFinish) {
        if (toLoad.length == 0) {
            appState.busy = false;
            ready();
            if (onFinish) {
                onFinish();
            }
            return;
        }
        let repo = toLoad.pop();
        loadRepo(repo)
            .then(result => pako.ungzip(result.archive))
            .then(arr => arr.buffer)
            .then(buf => untar(buf))
            .then(files => {
            var repoData = new RepoData();
            var dec = new TextDecoder("utf-8");
            for (let rawFile of files) {
                let f = new RepoFileData();
                f.name = rawFile.name;
                f.contents = dec.decode(rawFile.buffer);
                repoData.files.push(f);
            }
            if (repoData.files.length != repo.Files.length) {
                console.error(`${repo.Name}: unpacked files count mismatch: have ${repoData.files.length}, want ${repo.Files.length}`);
                if (repoData.files.length > repo.Files.length) {
                    let wantFiles = {};
                    for (let f of repo.Files) {
                        wantFiles[f.Name] = true;
                    }
                    let unexpectedFiles = [];
                    for (let f of repoData.files) {
                        let name = f.name.slice(repo.Name.length + 1);
                        if (!(name in wantFiles)) {
                            unexpectedFiles.push(name);
                        }
                    }
                    console.error(`${repo.Name}: unexpected files`, unexpectedFiles);
                }
                else {
                    let haveFiles = {};
                    for (let f of repoData.files) {
                        let name = f.name.slice(repo.Name.length + 1);
                        haveFiles[name] = true;
                    }
                    let missingFiles = [];
                    for (let f of repo.Files) {
                        if (!(f.Name in haveFiles)) {
                            missingFiles.push(f.Name);
                        }
                    }
                    console.error(`${repo.Name}: missing files`, missingFiles);
                }
            }
            else {
                // Files count is correct. Check the order.
                for (let i in repoData.files) {
                    let f1 = repoData.files[i];
                    let name = f1.name.slice(repo.Name.length + 1);
                    let f2 = repo.Files[i];
                    if (name !== f2.Name) {
                        console.error(`${repo.Name}: files[${i}] name mismatch: ${name} vs ${f2.Name}`);
                        break;
                    }
                }
            }
            console.log(`loaded ${repo.Name} repo`);
            appState.corpus.set(repo.Name, repoData);
            let $checkbox = (document.getElementById(`repository-${repo.Name}`));
            $checkbox.parentElement.classList.add('blue-text');
            loadRecursive(toLoad, onFinish);
        });
    }
    function getSelectedRepos() {
        var selected = [];
        for (let i = 0; i < appState.metadata.Repositories.length; i++) {
            let repo = appState.metadata.Repositories[i];
            let $checkbox = (document.getElementById(`repository-${repo.Name}`));
            if ($checkbox.checked) {
                selected.push(repo);
            }
        }
        return selected;
    }
    function forcedLoadRepositories(onFinish = function () { }) {
        appState.busy = true;
        let toLoad = [];
        for (let repo of getSelectedRepos()) {
            if (!appState.corpus.has(repo.Name)) {
                toLoad.push(repo);
            }
        }
        loadRecursive(toLoad, onFinish);
    }
    function loadRepositories() {
        if (appState.busy) {
            return;
        }
        forcedLoadRepositories();
    }
    function allReposSelected() {
        return getSelectedRepos().length === appState.metadata.Repositories.length;
    }
    function updateCheckallButton(unselect) {
        let $selectAll = document.getElementById('selectall-button');
        if (unselect) {
            $selectAll.innerText = 'Unselect all';
        }
        else {
            $selectAll.innerText = 'Select all';
        }
    }
    function setupUI() {
        var $corpusSelection = document.getElementById('corpus-selection');
        var selectionHTML = '';
        var buildSelectionTable = () => {
            for (let i = 0; i < appState.metadata.Repositories.length; i++) {
                selectionHTML += '<tr>';
                for (let j = 0; j < 6; j++) {
                    var index = (i * 6) + j;
                    if (index >= appState.metadata.Repositories.length) {
                        return;
                    }
                    var repo = appState.metadata.Repositories[index];
                    let sizeMB = (repo.MinifiedSize * 0.000001).toFixed(1);
                    let hint = `Tags: [${repo.Tags}], Size: ${sizeMB} MB, Files: ${repo.Files.length}, SLOC: ${repo.SLOC.toLocaleString()}`;
                    selectionHTML += `<td><label title="${hint}"><input id="repository-${repo.Name}" type="checkbox"> ${repo.Name}</label></td>`;
                }
                selectionHTML += '</tr>';
            }
        };
        buildSelectionTable();
        $corpusSelection.innerHTML = '<table style="width: 100%; display: table">' + selectionHTML + '</table>';
        let $loadRepos = document.getElementById('load-button');
        $loadRepos.onclick = function () {
            loadRepositories();
        };
        let $selectAll = document.getElementById('selectall-button');
        $selectAll.onclick = function () {
            let allSelected = allReposSelected();
            let value = !allSelected;
            for (let repo of appState.metadata.Repositories) {
                let $checkbox = (document.getElementById(`repository-${repo.Name}`));
                $checkbox.checked = value;
            }
            updateCheckallButton(value);
        };
        for (let repo of appState.metadata.Repositories) {
            let $checkbox = (document.getElementById(`repository-${repo.Name}`));
            $checkbox.onchange = function () {
                updateCheckallButton(allReposSelected());
            };
        }
        var $run = document.getElementById('run-button');
        var $results = document.getElementById('search-results');
        $results.innerHTML = '';
        $run.onclick = function () {
            if (appState.running) {
                appState.running = false;
                return;
            }
            if (appState.busy) {
                return;
            }
            appState.busy = true;
            let pattern = document.getElementById('search-pattern').value;
            let filter = document.getElementById('results-filter').value;
            $results.innerHTML = '';
            appState.searchResults.clear();
            appState.filesScanned = 0;
            appState.slocProcessed = 0;
            appState.filesTotal = 0;
            appState.hits = 0;
            let repos = getSelectedRepos();
            for (let repo of repos) {
                appState.filesTotal += repo.Files.length;
            }
            forcedLoadRepositories(() => {
                let $progress = document.getElementById('search-progress');
                $progress.innerHTML = '';
                $run.innerText = 'Stop';
                appState.running = true;
                appState.runStartTime = window.performance.now();
                runQueryRecursive(pattern, filter, repos);
            });
        };
    }
    function runApp(wasm) {
        ready();
        console.log("runApp()");
        appState.wasm = wasm;
        appState.go.run(wasm);
        setupUI();
    }
    function initPolyfills() {
        if (!WebAssembly.instantiateStreaming) {
            WebAssembly.instantiateStreaming = (resp, importObject) => __awaiter(this, void 0, void 0, function* () {
                const r = yield (yield resp).arrayBuffer();
                return yield WebAssembly.instantiate(r, importObject);
            });
        }
    }
    function main() {
        initPolyfills();
        loadMetadata().
            then(metadata => loadWASM(metadata)).
            then(wasm => runApp(wasm));
    }
    App.main = main;
})(App || (App = {}));
window.onload = function () {
    App.main();
};
