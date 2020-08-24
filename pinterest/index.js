let fs = require('fs');
const mark = '236x';

let express = require('express');
let app = express();

let bodyParser = require('body-parser');

app.use(require('cors')());

app.use(bodyParser.urlencoded({
  extended: true
}));
app.use(bodyParser.json())

app.post('/', (req, res, next) => {
  console.log('BODY', req.body);
  res.send('OK');
})

app.use(express.static(__dirname + '/browser'));

app.listen(3000, function () {
  console.log('listening...', 3000);
})

// let data = fs.readFileSync('./imageLinks.txt');
// let pics = data.toString().split('\n');
// let dup = {};
// pics.forEach(pic => dup[pic] = true);
// let result = Object.keys(dup);
// result = result.map(url => {
//   let i = url.indexOf(mark);
//   let first = url.substring(0, i);
//   let end = url.substring(i + mark.length);
//   let fin = first + 'originals' + end;
//   return fin;
// })

// fs.writeFileSync('./result.txt', result.join('\n'));

// let data = fs.readFileSync('./result.txt');
// let pics = data.toString().split('\n');
// // console.log(pics);

// fs.writeFileSync('./result2.txt', JSON.stringify(pics));

// let request = require('request');
// let axios = require('axios');

// var download = function (uri, filename, callback) {
//   request.head(uri, function (err, res, body) {
//     console.log('content-type:', res.headers['content-type']);
//     console.log('content-length:', res.headers['content-length']);

//     request(uri).pipe(fs.createWriteStream(filename)).on('close', callback);
//   });
// };

// let data = fs.readFileSync('./result.txt');
// let pics = data.toString().split('\n');
// let i = 0;
// pics.forEach(pic => {
// download(pics[28], './images/reminisce-' + ++i + '.png', function () { console.log('dun') });
// // })
// const download_image = (url, image_path) =>
//   axios({
//     url,
//     responseType: 'stream',
//     headers: { 'user-agent': "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36" }
//   }).then(
//     response =>
//       new Promise((resolve, reject) => {
//         response.data
//           .pipe(fs.createWriteStream(image_path))
//           .on('finish', () => resolve())
//           .on('error', e => reject(e));
//       }),
//   );

// download_image(pics[0], './images/reminisce-' + ++i)
//   .then(res => console.log('RES', res))
//   .catch(rej => console.error("ERR", rej))

// download('https://www.google.com/images/srpr/logo3w.png', 'google.png', function () {
//   console.log('done');
// });

// Remove duplicates from the original file
// let data = fs.readFileSync('./imageLinks.txt');
// let pics = data.toString().split('\n');
// let dup = {};
// pics.forEach(pic => {
//   if (pic.indexOf('75x75') < 0 && pic.indexOf('user') < 0) {
//     dup[pic] = true
//   }
// });
// console.log(Object.keys(dup).length, pics.length);

// fs.writeFileSync('./result.txt', Object.keys(dup).join('\n'));