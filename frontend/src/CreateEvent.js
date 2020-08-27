import React from "react";
import axios from "axios";
import { Form } from "react-bootstrap";

export default class CreateEvent extends React.Component {
  constructor() {
    super();
    this.state = {
      title: "",
      desc: "",
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
      .post(`/events`, {
        title: this.state.title,
        desc: this.state.desc,
      })
      .then((res) => {
        console.log(res.data);
      })
      .catch((error) => {
        console.log(error);
      });
    event.target.reset();
  }

  render() {
    return (
      <Form className="create-event-form" onSubmit={this.handleSubmit}>
        <div className="input">
          <Form.Label>Event Name</Form.Label>
          <Form.Control
            type="text"
            name="title"
            onChange={this.handleChange}
            placeholder="Enter Event Name"
          />
        </div>
        <div className="input">
          <Form.Label>Description</Form.Label>
          <Form.Control
            type="text"
            name="desc"
            onChange={this.handleChange}
            placeholder="Description"
          />
        </div>

        <button className="blue-button" type="submit">
          Add
        </button>
      </Form>
    );
  }
}
