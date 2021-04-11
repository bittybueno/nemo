import React, { useState } from "react";
import axios from "axios";
import { Form } from "react-bootstrap";

export function CreateEvent() {
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");

  const handleTitleChange = (e) => {
    setTitle(e.target.value);
  };

  const handleDescChange = (e) => {
    setDesc(e.target.value);
  };

  const handleSubmit = () => {
    axios
      .post(`/events`, {
        title: title,
        desc: desc,
      })
      .then((res) => {
        console.log(res.data);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <Form className="create-event-form" onSubmit={handleSubmit}>
      <div className="input">
        <Form.Label>Event Name</Form.Label>
        <Form.Control
          type="text"
          name="title"
          onChange={handleTitleChange}
          placeholder="Enter Event Name"
        />
      </div>
      <div className="input">
        <Form.Label>Description</Form.Label>
        <Form.Control
          type="text"
          name="desc"
          onChange={handleDescChange}
          placeholder="Description"
        />
      </div>

      <button className="blue-button" type="submit">
        Add
      </button>
    </Form>
  );
}
