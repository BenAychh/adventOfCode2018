const fs = require('fs');

const file = process.argv[2];
const jsonName = `${file.split('.data')[0]}.json`;

const contents = fs.readFileSync(file, 'utf8');
const contentsAsArray = contents.split('\n').filter(string => string !== '');

fs.writeFileSync(jsonName, JSON.stringify(contentsAsArray));