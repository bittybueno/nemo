import React from "react";
import axios from "axios";

import { Form } from "react-bootstrap";

export default class CreateQuestion extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      body: "",
      update: false,
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({
      [event.target.name]: event.target.value,
    });
  }

  handleSubmit(event) {
    axios
      .post(`/questions`, {
        event_id: this.props.eventID,
        body: this.state.body,
      })
      .then((res) => {
        console.log(res.data);
      })
      .catch((error) => {
        console.log(error);
      });
    event.target.reset();
    this.props.cb();
  }

  render() {
    return (
      <div>
        <Form
          className="create-question-form"
          id="myForm"
          onSubmit={this.handleSubmit}
        >
          <div className="input">
            <Form.Label>Enter Question</Form.Label>
            <Form.Control
              type="text"
              name="body"
              onChange={this.handleChange}
              placeholder="Question"
            />
          </div>

          <button className="blue-button" type="submit">
            Add
          </button>
        </Form>
      </div>
    );
  }
}
