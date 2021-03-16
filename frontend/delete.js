async function deleteISBN(isbns) {

    // Check out the given ISBNs.
    return swaggerClient
        .then(
            client => client.apis.api.bookDelete({isbns: isbns}),
            reason => console.error('failed to load the spec: ' + reason)
        )
        .then(
            bookDeleteResult => { /* No operation.*/
            },
            reason => console.error('failed api call: ' + reason)
        );
}
