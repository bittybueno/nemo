import React from "react";
import axios from "axios";
import { Link } from "react-router-dom";

class Events extends React.Component {
  state = {
    events: [],
  };

  componentDidMount() {
    axios
      .get(`/events`)
      .then((res) => {
        console.log(res.data);
        const events = res.data;
        this.setState({ events });
      })
      .catch((error) => {
        console.log(error);
      });
  }

  render() {
    return (
      <ul className="list-events">
        {this.state.events.map((event) => (
          <li key={event.id}>
            <Link
              style={{ textDecoration: "none" }}
              className="link"
              to={`/events/${event.id}`}
            >
              {event.title}
            </Link>
          </li>
        ))}
      </ul>
    );
  }
}

export default Events;
