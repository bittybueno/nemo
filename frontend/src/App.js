import React, { Component } from "react";
import "./App.css";
import Events from "./Events.js";
import Header from "./Header.js";
import CreateEvent from "./CreateEvent.js";

import { Link } from "react-router-dom";

class App extends Component {
  render() {
    return (
      <div className="App">
        <Header></Header>
        <CreateEvent className="create-event-form"></CreateEvent>
        <Events className="list-events-background"></Events>
        <Link to="/question-board">
          <button className="blue-button">Go to Questions</button>
        </Link>
      </div>
    );
  }
}

export default App;
