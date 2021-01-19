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
  ).innerHTML += `<li><div class="box ${language}"></div>  ${language}</li>`)
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

function populateStats(node, parent = "statsList") {
  if (!node.hasOwnProperty("children")) {
    document.getElementById(parent).innerHTML += `<li>${node.name}
<div style="display: float; float: right; padding-right: 2em; color: ${node.color}; font-weight: bold">${node.size}</div></li>`;
    console.log(node.name, node.size);
  } 
  else {
    let nestedId = `li_${node.name}`;
    document.getElementById(parent).innerHTML += `
<li>/<b>${node.name}</b></li>
  <ul id=${nestedId}>
`;
    node.children.map((node) => populateStats(node, nestedId));
    document.getElementById(parent).innerHTML += `</ul>`;
  }
}

wrapper("statsList", data);
