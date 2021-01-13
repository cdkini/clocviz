/* Inject CSS: https://stackoverflow.com/questions/15505225/inject-css-stylesheet-as-string-using-javascript */
function addStyle(styleString) {
  const style = document.createElement('style');
  style.textContent = styleString;
  document.head.append(style);
}

const languages = data.children.map(child => child.name);


languages.forEach(language => 
    document.getElementById('legend').innerHTML +=
        `<li><div class="box ${language}"></div>${language}</li>`
);

data.children.forEach(child => 
    addStyle(`
  .box.${child.name} {
    background: ${child.color};
  }
  `)
);





