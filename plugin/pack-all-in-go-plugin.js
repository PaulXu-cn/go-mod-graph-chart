class PackAllInGoPlugin {
  apply(compiler) {
    // emit is asynchronous hook, tapping into it using tapAsync, you can use tapPromise/tap(synchronous) as well
    compiler.hooks.emit.tapAsync('PackAllInGoPlugin', (compilation, callback) => {
      // Create a header string for the generated file:
      var filelist = '';
      var indexHtml, mainJs;
      var goCode = `package godist

func GetFile(file string) string {
  switch {
  case \`index.html\` == file:
    return IndexHtml
  case \`main.js\` == file:
    return MainJson
  case \`favicon.ico\` == file:
    return Favicon
  default:
    return ""
  }
}

var IndexHtml = \`--index.html--\`

var MainJson = \`--main.js--\`

var Favicon = \`favicon.ico\``

      // Loop through all compiled assets,
      // adding a new line item for each filename.
      for (var filename in compilation.assets) {
        if ("main.js" == filename) {
          let jsCode = compilation.assets[filename].source()
          let jsCodeString = jsCode.slice();
          jsCodeString = jsCodeString.replace(/\`/g, "\` + \"\`\" + \`")
          goCode = goCode.replace('--main.js--', jsCodeString)
        } else if ("index.html") {
          let htmlCode = compilation.assets[filename].source()
          goCode = goCode.replace('--index.html--', htmlCode)
        }
      }
  
      // 将这个列表作为一个新的文件资源，插入到 webpack 构建中：
      compilation.assets['../godist/static.go'] = {
        source: function() {
          return goCode;
        },
        size: function() {
          return goCode.length;
        }
      };

      callback();
    });
  }
}

module.exports = PackAllInGoPlugin;
