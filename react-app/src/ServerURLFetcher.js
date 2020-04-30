
function ServerURLFetcher() {
    this.url = "http://localhost:8080"
}

ServerURLFetcher.prototype.getURL = function() {
    return this.url;
}

export default new ServerURLFetcher();