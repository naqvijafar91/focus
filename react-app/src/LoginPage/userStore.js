function Store(){
    // Load user from localstorage and set value of isAuthenticated
    this.key= "focus-app-user";

    if(JSON.parse(localStorage.getItem("focus-app-user"))) {
        
        this.user=JSON.parse(localStorage.getItem("focus-app-user"));
        this.isAuthenticated=true;
        // debugger;
    }

}

Store.prototype.deleteUser = function() {
    localStorage.removeItem(this.key);
};

Store.prototype.saveUser=function(user) {
    this.user=user;
    localStorage.setItem("focus-app-user", JSON.stringify(this.user));
    // console.log(this)
    this.isAuthenticated=true;
}

Store.prototype.getUser=function(){
    return this.user;
}
export default new Store();