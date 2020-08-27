import React from "react";
import axios from "axios";

import Questions from "./Questions.js";
import CreateQuestion from "./CreateQuestion.js";
import Event from "./Event.js";

export default class QuestionBoard extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      event_id: this.props.event_id,
      questions: [],
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleClick = this.handleClick.bind(this);
    this.fetchData = this.fetchData.bind(this);
  }

  fetchData = () => {
    axios
      .get(`/questions`, {
        params: {
          event_id: this.state.event_id,
        },
      })
      .then((res) => {
        console.log(res.data);
        const questions = res.data;
        this.setState({ questions: questions });
      })
      .catch((error) => {
        console.log(error);
      });
  };

  handleChange(event) {
    this.setState({
      event_id: event.target.value,
    });
  }

  handleClick(event) {
    event.preventDefault();
    this.setState({ isSubmitted: true });
    this.fetchData();
  }

  render() {
    return (
      <div>
        <Event cb={this.fetchData} eventID={this.state.event_id}></Event>
        <Questions
          questions={this.state.questions}
          eventID={this.state.event_id}
        ></Questions>
        <CreateQuestion
          cb={this.fetchData}
          eventID={this.state.event_id}
        ></CreateQuestion>
      </div>
    );
  }
}
