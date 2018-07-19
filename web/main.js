// var IdA = document.getElementsByClassName('Id');
// var KisimA = document.getElementsByClassName('Kisim');
// var KucretA = document.getElementsByClassName('Kucret');
// var KyayinA = document.getElementsByClassName('Kyayin');
// var deletes = document.getElementsByClassName('remove');
// var updates = document.getElementsByClassName('update');
//
// var btnGet = document.getElementById('Rget');
// var btnPost = document.getElementById('Rpost');
// var btnDelete = document.getElementById('Rdelete');
// var btnPut = document.getElementById('Rput');

var url = "http://localhost:2550/book"
var book = {};

var xhr  = new XMLHttpRequest()
xhr.open('GET', url + '/' + 1, true)
xhr.onload = function () {
	book = JSON.parse(xhr.responseText);
  console.log("Get edilen book:", book);
	if (xhr.readyState == 4 && xhr.status == "200") {
    update(book);
		// console.table(book);
	} else {
		console.error(book);
	}
}
xhr.send(null);

// UPDATE [ PUT ]

function update(book) {
  book.Kisim  = "Snow2";
  var json = JSON.stringify(book);

  var req = new XMLHttpRequest();
  req.open("PUT", url+'/1', true);

  req.onload = function () {
    var updatedBook = JSON.parse(req.responseText);
    if (req.readyState == 4 && req.status == "200") {
      console.log("Update edilen book:", updatedBook);
      listBooks();
    } else {
      console.error(updatedBook);
    }
  }
  req.send(json);
}

// LİSTELE [ GET ]

function listBooks() {
  var lst = new XMLHttpRequest()
  lst.open('GET', url, true)

  lst.onload = function () {
  	var lstbooks = JSON.parse(lst.responseText);
    if (lst.readyState == 4 && lst.status == "200") {
      console.log("Books: ", lstbooks);
      deleteBooks();
    } else {
      console.error(lstbooks);
    }
  }
  lst.send(null);
}

// DELETE [ DELETE ]

function deleteBooks() {
  var dlt = new XMLHttpRequest()
  dlt.open("DELETE", url + "/3", true)

  dlt.onload = function () {
    // var deletebooks = JSON.parse(dlt.responseText);
    if (dlt.readyState == 4 && dlt.status == "200") {
      // console.log("Delete edilen book: ", deletebooks);
      addBooks();
    } else {
      console.error(deletebooks);
    }
  }
  dlt.send(null);
}

// ADD [ POST ]

function addBooks() {
  var add = new XMLHttpRequest();
  book.Kisim = "Yeni Kitap"
  book.Kucret = 50;
  book.Kyayin = 2018

  var json = JSON.stringify(book);
  var addy = new XMLHttpRequest();
  addy.open("POST", url + "/4", true);
  addy.setRequestHeader('Content-type','application/json; charset=utf-8');
  addy.onload = function () {
    var addBook = JSON.parse(addy.responseText);

    if (addy.readyState == 4 && addy.status == "200") {
      console.log("Add edilen book:", addBook);
    } else {
      console.error(updatedBook);
    }
  }
  addy.send(json);
}
