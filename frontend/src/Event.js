import React from "react";
import axios from "axios";

import Jumbotron from "react-bootstrap/Jumbotron";

class Event extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      eventTitle: "",
      eventDesc: "",
    };
  }

  componentDidMount() {
    axios
      .get(`/events`, {
        params: {
          event_id: this.props.eventID,
        },
      })
      .then((res) => {
        console.log(res.data);
        this.setState({
          eventTitle: res.data.title,
          eventDesc: res.data.desc,
        });
      })
      .catch((error) => {
        console.log(error);
      });
    this.props.cb();
  }

  render() {
    return (
      <div className="jumbo">
        <Jumbotron fluid>
          <h1>{this.state.eventTitle}</h1>
          <p>{this.state.eventDesc}</p>
        </Jumbotron>
      </div>
    );
  }
}

export default Event;
