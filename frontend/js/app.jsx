class App extends React.Component {
    render() {
        return (<Home />);

    }
  }

class Home extends React.Component {
render() {
    return (
    <div className="container">
            <h1>morethanjustlinks</h1><form action="/login" method="post" onsubmit="submitForm(event)">
            <div>Mobile/Username/Email <input type="text" name="name" id="name" required /></div>
            <div>Password <input type="password" name="psword" id="psword" required /></div>
            <div><input type="submit" value="Login" /></div>
            <div><a href="/newAccountForm">Create a new account</a> | <a href="#">Forgot password?</a> </div>
        </form><footer>Copyright &copy; 2023 yourUserName All Rights Reserved</footer>
    </div>
    )
}
}

ReactDOM.render(<App />, document.getElementById('app'));
