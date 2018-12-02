const fs = require('fs')

function loadData(fileName) {
  return new Promise((resolve, reject) => {
    fs.readFile(fileName, 'utf8', (err, data) => {
      if (err) {
        reject(err)
      }
      resolve(data.split('\n'));
    });
  })
}

module.exports = {
  loadData
}