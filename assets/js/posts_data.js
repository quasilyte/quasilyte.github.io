"use strict";

(function() {
  // Common categories
  var ragf = "Ragf project";

  // Possible tags 
  var tagList = [
    "DSL",
    "x86_64",
    "assembly",
    "nasm",
    "perfomance",
    "Ruby"
  ];

  var posts = [
    {
      "category": ragf,
      "title": "Red assembly goez fasta!",
      "date": 1475996682,
      "file": "ragf",
      "tags": [
        "DSL", "Ruby"
      ]
    },
    {
      "category": ragf,
      "title": "nbody in x86_64 assembly",
      "date": 1475844010,
      "file": "nbody",
      "tags": [
        "x86_64",
        "assembly",
        "nasm",
        "perfomance",
      ],
      "hidden": true
    },
    {
      "category": "Ragf project",
      "title": "Dummy entry1",
      "date": 1475844050,
      "file": "",
      "tags": []
    },
    {
      "category": "Ragf project",
      "title": "Dummy entry2",
      "date": 1475844050,
      "file": "",
      "tags": []
    }
  ];

  // Adding id to each post 
  _.each(posts, (post, id) => post.id = id); 

  // Tag validation
  _.each(App.posts, function(post) {
    _.each(post.tags, function(tag) {
      if (!_.contains(tagList, tag)) {
        throw `tag list has no tag "${tag}"`;
      }
    })
  });

  App.rodata.tagList = tagList; 
  App.posts = posts; 
}());
