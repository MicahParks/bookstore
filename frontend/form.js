// Assign some elements to the outer scope. Not great but convenient.
let authorElem = document.getElementById("bookAuthor");
let descriptionElem = document.getElementById("bookDescription");
let isbnElem = document.getElementById("bookIsbn");
let titleElem = document.getElementById("bookTitle");
let quantityElem = document.getElementById("bookQuantity");
let quantityDiv = document.getElementById("quantityDiv");

function clearForm() {

    // Populate the form data.
    authorElem.value = "";
    descriptionElem.value = "";
    isbnElem.value = "";
    titleElem.value = "";
    quantityDiv.style.visibility = 'visible';
    quantityElem.value = "1";
}

function populateForm(book) {

    // Populate the form data.
    authorElem.value = book.author;
    descriptionElem.value = book.description;
    isbnElem.value = book.isbn;
    titleElem.value = book.title;
    quantityDiv.style.visibility = 'hidden';
}

async function submitForm(e) {
    e.preventDefault();

    // Disable the submit button from being hit more than once.
    const submitButton = document.getElementById("submitButton");
    submitButton.disabled = true;
    setTimeout(() => submitButton.disabled = false, 1000);

    // Create the book from the form values.
    let book = {
        author: authorElem.value,
        description: descriptionElem.value,
        isbn: isbnElem.value,
        title: titleElem.value
    };

    let bookQuantities = {};
    bookQuantities[book.isbn] = {
        book: book,
        quantity: Number(quantityElem.value),
    };

    // Write the books.
    writeBooks(bookQuantities).then(function () {

        // Repopulate the table.
        buildTable();

        // Hide the modal.
        $("#formModal").modal("hide");
    });
}
