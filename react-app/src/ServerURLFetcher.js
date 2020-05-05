
function ServerURLFetcher() {
    this.url_dev = "http://localhost:8080";
    this.url_prod = "http://localhost:8080";
}

ServerURLFetcher.prototype.getURL = function() {
    if (process.env.REACT_APP_ENV === 'prod') {
        console.log("This is prod");
        return this.url_prod;
    }
    return this.url_dev;
}

export default new ServerURLFetcher();