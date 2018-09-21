import React, { Component } from 'react';
import './App.css';
import GoogleLogin from 'react-google-login';
import Client from './Client';

class App extends Component {

  constructor(props) {
    super(props)

    this.responseGoogle = this.responseGoogle.bind(this);
  }

  responseGoogle(response) {
    console.log(response);
    alert(Client.getHealth(response.idToken));
  }

  loginPage() {
    return (
      <div className='App'>
        <div className='ui text container'>
        <GoogleLogin
           clientId="217209923893-kv53i3hqgk1plrapk0eub1p4jr7sipet.apps.googleusercontent.com"
           buttonText="Sign in with Google"
           onSuccess={this.responseGoogle}
           onFailure={this.responseGoogle}
         />
        </div>
      </div>
    );
  }

  render() {
    return this.loginPage();
  }
}

export default App;
