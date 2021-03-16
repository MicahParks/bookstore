async function readBooks(isbns) {

    // Get the Book data for the ISBNs.
    let books;
    let promise = swaggerClient
        .then(
            client => client.apis.api.bookRead({isbns: isbns}),
            reason => console.error('failed to load the spec: ' + reason)
        )
        .then(
            bookReadResult => books = JSON.parse(bookReadResult.data),
            reason => console.error('failed api call: ' + reason)
        );
    await promise;

    return books;
}

async function writeBooks(books) {

    // Write the book data.
    return swaggerClient
        .then(
            client => client.apis.api.bookWrite({books: books, operation: "upsert"}),
            reason => console.error('failed to load the spec: ' + reason)
        )
        .then(
            bookWriteResult => { /* No operation.*/
            },
            reason => console.error('failed api call: ' + reason)
        );
}
