
function ServerURLFetcher() {
    this.url = "http://localhost:8080";
    if(process.env.REACT_APP_URL) {
        this.url = process.env.REACT_APP_URL;
    }
}

ServerURLFetcher.prototype.getURL = function() {
    return this.url;
}

export default new ServerURLFetcher();