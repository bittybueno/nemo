import React from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import QuestionBoard from "./QuestionBoard";

class LinkToQuestions extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      questions: [],
    };
  }

  componentDidMount() {
    console.log("here");
    axios
      .get(`/events?`, {
        params: {
          event_id: this.props.match.params.id,
        },
      })
      .then((res) => {
        const events = res.data;
        console.log(res);
        this.setState({ events });
      })
      .catch((error) => {
        console.log(error);
      });
  }

  render() {
    return (
      <div>
        <Link to="/">
          <button>Back</button>
        </Link>
        <QuestionBoard event_id={this.props.match.params.id}></QuestionBoard>
      </div>
    );
  }
}

export default LinkToQuestions;
