/* Inject CSS: https://stackoverflow.com/questions/15505225/inject-css-stylesheet-as-string-using-javascript */
function addStyle(styleString) {
  const style = document.createElement("style");
  style.textContent = styleString;
  document.head.append(style);
}

const languages = data.children.map((child) => child.name);

languages.forEach(
  (language) =>
    (document.getElementById(
      "legend"
    ).innerHTML += `<li><div class="box ${language}"></div>${language}</li>`)
);

data.children.forEach((child) =>
  addStyle(`
  .box.${child.name} {
    background: ${child.color};
  }
  `)
);

function wrapper(id, node) {
  clear(id);
  populateStats(node);
}

function clear(id) {
  document.getElementById(id).innerHTML = "";
}

function populateStats(node) {
  if (!node.hasOwnProperty("children")) {
    document.getElementById(
      "statsList"
    ).innerHTML += `<li>${node.name}: ${node.size}</li>`;
    console.log(node.name, node.size);
  } else {
    console.log("Folder", node.name);
    node.children.map(populateStats);
  }
}

wrapper("statsList", data);
