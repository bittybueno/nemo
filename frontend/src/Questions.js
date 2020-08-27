import React from "react";
import { Card, CardColumns } from "react-bootstrap";

import "./Card.css";

class Questions extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      questions: this.props.questions,
      display: false,
    };
    this.renderTableData.bind(this);
  }

  renderTableData() {
    return this.props.questions.map((question) => {
      return (
        <Card className="question-card" key={question.id}>
          <Card.Body>{question.body}</Card.Body>
        </Card>
      );
    });
  }

  render() {
    return (
      <div>
        <CardColumns>
          {this.props.questions !== null ? this.renderTableData() : <p>none</p>}
        </CardColumns>
      </div>
    );
  }
}
export default Questions;
