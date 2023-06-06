// Many code fragments are based on http://webglplayground.net/.

function main() {
  let renderingContext = {
    "width": 100,
    "height": 100,
    "maxWidth": 0,
    "maxHeight": 0,
    "resized": false,
  };

  const FPS = 60.0;
  const FRAME_DELTA = 1.0 / FPS;

  const renderer = new THREE.WebGLRenderer({
    // "antialias": true,
  });
  const gl = renderer.getContext();
  const scene = new THREE.Scene();
  let sceneObject = null;
  const camera = new THREE.OrthographicCamera(-1, 1, 1, -1, 0.1, 100.0);
  camera.position.z = 1.0;
  const uniforms = {
    "cursorPos": {"value": new THREE.Vector2(0.0, 0.0)},
  };
  const textureLoader = new THREE.TextureLoader();
  let currentTextureIndex = 0;
  let editingCode = true;
  let currentFragmentShader = '';
  const snippets = [];
  let cursorPos = {"x": 0.0, "y": 0.0};

  const $shaderError = document.getElementById("shader_error");
  const $editorContainer = document.getElementById("shader_editor_container");
  const $rendererContainer = document.getElementById("renderer_container");
  const $editor = ace.edit("shader_editor");
  const $textureSelect = document.getElementById("texture_select");
  const $snippetSelect = document.getElementById("snippet_select");
  const $blendingModeSelect = document.getElementById("blending_mode_select");
  const $renderButton = document.getElementById("render_button");
  const $shaderOptions = document.getElementById("shader_options");

  let textureFileInputs = [];
  for (let i = 1; i <= 3; i++) {
    textureFileInputs.push(document.getElementById(`texture_file${i}`))
  }
  let textureURLInputs = [];
  for (let i = 1; i <= 3; i++) {
    textureURLInputs.push(document.getElementById(`texture_url${i}`))
  }

  function replaceMesh() {
    if (sceneObject) {
      scene.remove(sceneObject);
    }

    const vertShader = `
      varying vec2 texPos;
      varying vec2 pos;
      uniform vec2 textureSize;

      void main() {
        vec4 mvPosition = modelViewMatrix * vec4(position, 1.0);
        gl_Position = projectionMatrix * mvPosition;

        texPos = uv;
        pos = vec2(textureSize.x * texPos.x, textureSize.y - (textureSize.y * texPos.y));
      }
    `;

    const materialOptions = {
      "uniforms": uniforms,
      "fragmentShader": currentFragmentShader,
      "vertexShader": vertShader,
    };
    const blendingMode = $blendingModeSelect.value;
    console.log(blendingMode);
    if (blendingMode === "overwrite") {
      materialOptions.blending = THREE.CustomBlending;
      materialOptions.blendEquation = THREE.AddEquation;
      materialOptions.blendSrc = THREE.SrcAlphaFactor;
      materialOptions.blendDst = THREE.OneMinusDstAlphaFactor;
    } else {
      materialOptions.blending = {
        "none": THREE.NoBlending,
        "normal": THREE.NormalBlending,
        "additive": THREE.AdditiveBlending,
        "subtractive": THREE.SubtractiveBlending,
        "multiply": THREE.MultiplyBlending,
      }[blendingMode];
    }

    const material = new THREE.ShaderMaterial(materialOptions);

    sceneObject = new THREE.Mesh(new THREE.BoxGeometry(2, 2), material);
    scene.add(sceneObject);
  }

  function updateShader(shaderText) {
    if (shaderText.trim() === '') {
      return 'Empty shader program, missing main() function';
    }
    if (shaderText === currentFragmentShader) {
      return '';
    }

    const SHADER_HEADER = `#ifdef GL_ES
precision highp float;
#endif
vec2 toTex(vec2 textureSize, vec2 pos) { return vec2(pos.x, textureSize.y-pos.y)/textureSize; }
vec2 fromTex(vec2 textureSize, vec2 texPos) { return textureSize*texPos; }
    `;
    const numShaderHeaderLines = SHADER_HEADER.split("\n").length;
    shaderText = SHADER_HEADER + shaderText;

    const shaderID = gl.createShader(gl.FRAGMENT_SHADER);
    gl.shaderSource(shaderID, shaderText);
    gl.compileShader(shaderID);
    const shaderCompiled = gl.getShaderParameter(shaderID, gl.COMPILE_STATUS);
    if (!shaderCompiled) {
      const compilationLogs = gl.getShaderInfoLog(shaderID);
      return compilationLogs.replaceAll(/ERROR: 0:(\d+):/g, function (match, capture) {
        return 'ERROR: line ' + (capture-numShaderHeaderLines+1) + ':';
      });
    }

    currentFragmentShader = shaderText;
    replaceMesh();
    return '';
  }

  function updateTexture(source, index) {
    const texture = textureLoader.load(source, function (t) {
      t.minFilter = THREE.NearestFilter;
      t.magFilter = THREE.NearestFilter;
      if (index === 0) {
        renderingContext.resized = t.image.width > renderingContext.maxWidth || t.image.height > renderingContext.maxHeight;
        renderingContext.width = t.image.width <= renderingContext.maxWidth ? t.image.width : renderingContext.maxWidth;
        renderingContext.height = t.image.height <= renderingContext.maxHeight ? t.image.height : renderingContext.maxHeight;
        renderer.setSize(renderingContext.width, renderingContext.height);
      }
      uniforms.textureWidth = {"value": renderingContext.width+0.0};
      uniforms.textureHeight = {"value": renderingContext.height+0.0};
      uniforms.textureSize = {"value": new THREE.Vector2(renderingContext.width, renderingContext.height)};
      if (index === 0) {
        uniforms.texture1 = {"value": texture};
      } else if (index === 1) {
        uniforms.texture2 = {"value": texture};
      } else {
        uniforms.texture3 = {"value": texture};
      }
    });
  }

  function renderFrame() {
    if (editingCode) {
      return;
    }
    uniforms.cursorPos.value = new THREE.Vector2(cursorPos.x, renderingContext.maxHeight-cursorPos.y);
    uniforms.time.value += FRAME_DELTA;
    renderer.render(scene, camera);
  }

  function updateTextureFromURL(url, textureIndex) {
    const img = new Image();
    img.src = url;
    img.setAttribute("crossOrigin", "anonymous");
    img.onload = function() {
      var memCanvas = document.createElement("canvas");
      memCanvas.width = this.width;
      memCanvas.height = this.height;
      var ctx = memCanvas.getContext("2d");
      ctx.drawImage(this, 0, 0);
      var dataURL = memCanvas.toDataURL("image/png");
      updateTexture(dataURL, textureIndex);
    };
  }

  function updateTextureFromFile(f, textureIndex) {
    let reader = new FileReader();
    reader.onload = (function(e) {
      updateTexture(e.target.result, textureIndex);
    });
    reader.readAsDataURL(f);
  }

  function onFileUpload(e) {
    let files = e.target.files;
    updateTextureFromFile(files[0], currentTextureIndex);
  }

  function onRenderPressed() {
    if (editingCode) {
      $renderButton.innerText = "Edit GLSL";
      const compilationLog = updateShader($editor.getValue());
      $editorContainer.hidden = true;
      $shaderOptions.hidden = true;
      $shaderError.hidden = !compilationLog;
      $rendererContainer.hidden = compilationLog;
      if (compilationLog) {
        $shaderError.innerText = compilationLog;
      }
    } else {
      $shaderError.hidden = true;
      $rendererContainer.hidden = true;
      $shaderOptions.hidden = false;
      $editorContainer.hidden = false;
      $renderButton.innerText = "Render";
    }

    editingCode = !editingCode;
  }

  function selectSnippet(i) {
    const index = +i;
    const s = snippets[index];
    $editor.setValue(s.text, 1);
    $blendingModeSelect.value = s.blend;
    if ("url1" in s) {
      textureURLInputs[0].value = s.url1;
      updateTextureFromURL(s.url1, 0);
    }
    if ("url2" in s) {
      textureURLInputs[1].value = s.url2;
      updateTextureFromURL(s.url2, 1);
    }
    if ("url3" in s) {
      textureURLInputs[2].value = s.url3;
      updateTextureFromURL(s.url3, 2);
    }
  }

  function initPage() {
    const $editorDiv = document.getElementById("shader_editor");
    renderingContext.maxWidth = $editorDiv.offsetWidth;
    renderingContext.maxHeight = $editorDiv.offsetHeight;

    $rendererContainer.addEventListener("mousemove", (e) => {
      if (!e) {
        return;
      }
      const rect = e.target.getBoundingClientRect();
      cursorPos.x = e.clientX - rect.left;
      cursorPos.y = e.clientY - rect.top;
    });

    for (let $input of textureFileInputs) {
      $input.addEventListener("change", onFileUpload);
    }
    for (let $input of textureURLInputs) {
      $input.setAttribute("title", "An external image URL (e.g. imgur-hosted image), press enter to load");
      $input.addEventListener("keydown", function (e) {
        if (e.key === "Enter") {
          updateTextureFromURL(textureURLInputs[currentTextureIndex].value);
        }
      });
    }

    $renderButton.addEventListener("click", onRenderPressed);

    $textureSelect.addEventListener("change", function() {
      textureFileInputs[currentTextureIndex].hidden = true;
      textureURLInputs[currentTextureIndex].hidden = true;
      currentTextureIndex = +this.value;
      textureFileInputs[currentTextureIndex].hidden = false;
      textureURLInputs[currentTextureIndex].hidden = false;
    });

    const NUM_SNIPPETS = 7;
    for (let i = 0; i < NUM_SNIPPETS; i++) {
      const $elem = document.getElementById(`snippet${i}`);
      snippets.push({
        "text": $elem.innerText.trim() + "\n",
        "name": $elem.getAttribute("data-name"),
        "url1": $elem.getAttribute("data-url1"),
        "url2": $elem.getAttribute("data-url2"),
        "url3": $elem.getAttribute("data-url3"),
        "blend": $elem.getAttribute("data-blend"),
      });
    }
    const rawParts = [];
    for (let i in snippets) {
      const s = snippets[i];
      const raw = `<option value="${i}">${s.name}</option>`;
      rawParts.push(raw);
    }
    $snippetSelect.innerHTML = rawParts.join("\n");

    $snippetSelect.addEventListener('change', function() {
      selectSnippet(this.value);
    });

    $blendingModeSelect.addEventListener('change', function() {
      replaceMesh();
    });
  }

  function initEditor() {
    $editor.setOptions({
      "fontFamily": "mono",
      "fontSize": "14pt",
      "showFoldWidgets": false,
    });
    $editor.setTheme("ace/theme/monokai");
    $editor.session.setMode("ace/mode/glsl");
    selectSnippet(0);
  }

  function initScene() {
    renderer.setSize(renderingContext.width, renderingContext.height);
    $rendererContainer.appendChild(renderer.domElement);

    scene.add(camera);

    for (let i in textureFileInputs) {
      let $input = textureFileInputs[i];
      if ($input.files.length > 0) {
        console.log(`loaded file[${i}]`);
        updateTextureFromFile($input.files[0], +i);
      }
    }

    uniforms.time = {"type": "f", "value": 0};
  }

  initPage();
  initEditor();
  initScene();
  setInterval(renderFrame, 1000.0 / FPS);
  console.log("initialization completed");
}

window.addEventListener("load", function () {
  main();
});
