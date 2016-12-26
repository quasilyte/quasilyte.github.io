"use strict";

(function() {
  // Common categories
  var ragf = "Ragf project";
  var mixed = "Mixed";

  // Possible tags 
  var tagList = [
    "Ruby lang",
    "C lang",
    "C11",

    "Language design",
		
    "DSL",
    "x86_64",
    "assembly",
    "nasm",
    "perfomance",
    
    "hacks",
		"rants"
  ];

  var posts = [
    {
      "category": ragf,
      "title": "Red assembly goez fasta!",
      "date": 1475996682,
      "file": "ragf",
      "tags": [
        "DSL", 
        "Ruby lang"
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
      "category": mixed,
      "title": "C APIs and closures",
      "date": 1480146077,
      "file": "c_closures",
      "tags": [
        "C lang",
        "C11",
        "hacks"
      ]
    },
		{
			"category": mixed,
			"title": "Broken C defaults",
			"date": 1482764456,
			"file": "c_rants",
			"tags": [
				"C lang",
				"rants",
				"Language design",
			]
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
