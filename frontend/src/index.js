import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter, Route, Switch } from "react-router-dom";

import App from "./App";
import QuestionBoard from "./QuestionBoard";
import LinkToQuestions from "./LinkToQuestions";

import "./index.css";
import "bootstrap/dist/css/bootstrap.css";

ReactDOM.render(
  <BrowserRouter>
    <Switch>
      <Route path="/events/:id" component={LinkToQuestions} />
      <Route exact path="/" component={App} />
      <Route path="/question-board" component={QuestionBoard} />
    </Switch>
  </BrowserRouter>,
  document.getElementById("root")
);
