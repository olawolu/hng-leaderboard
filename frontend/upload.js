// var getJSON = function(callback) {
//     var url = "http://localhost:8080/upload"
//     var xhr = new XMLHttpRequest();
//     xhr.open('GET', url, true);
//     xhr.responseType = 'json';
//     xhr.onload = function() {
//       var status = xhr.status;
//       if (status === 200) {
//         callback(null, xhr.response);
//       } else {
//         callback(status, xhr.response);
//       }
//     };
//     xhr.send();
// };

// function getJSON() {
//   var url = "http://localhost:8080/upload";
//   var xhr = new XMLHttpRequest();
//   xhr.open("GET", url, true);
//   xhr.responseType = "json";
//   xhr.onload = function () {
//     var status = xhr.status;
//     if (status === 200) {
//       callback(null, xhr.response);
//     } else {
//       callback(status, xhr.response);
//     }
//   };
//   xhr.send();
// }

function handleUpload() {
  event.preventDefault();
  let file = document.getElementById("file").files[0];
  // const files = event.target.files
  let formData = new FormData();
  const url = "http://localhost:8080/upload";

  formData.append("file", file);

  fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => response.text())
    .then((data) => {
      console.log(typeof data);
      for (element in data) {
        console.log(element.Name);
      }
    })
    .catch((error) => {
      console.error(error);
    });
}

// document.getElementById('file').addEventListener('change', event => {
//     handleUpload(event)
// })
