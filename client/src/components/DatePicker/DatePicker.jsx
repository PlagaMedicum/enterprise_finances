import React from "react";
import FormControl from "react-bootstrap/FormControl";
import InputGroup from "react-bootstrap/InputGroup";

class DatePicker extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      day: 0,
      month: 0,
      year: 0
    }
  }

  render() {
    return (
        <InputGroup>
          <FormControl placeholder="day" onChange={e => this.props.setters.setDay(e.target.value)}/>
          <select className="custom-select" onChange={e => this.props.setters.setMonth(e.target.value)}>
            <option value="1">jan</option>
            <option value="2">feb</option>
            <option value="3">mar</option>
            <option value="4">apr</option>
            <option value="5">may</option>
            <option value="6">jun</option>
            <option value="7">jul</option>
            <option value="8">aug</option>
            <option value="9">sep</option>
            <option value="10">oct</option>
            <option value="11">nov</option>
            <option value="12">dec</option>
          </select>
          <FormControl placeholder="year" onChange={e => this.props.setters.setYear(e.target.value)}/>
        </InputGroup>
    );
  }
}

export default DatePicker;
