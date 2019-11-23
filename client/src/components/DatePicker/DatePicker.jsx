import React from "react";
import FormControl from "react-bootstrap/FormControl";
import InputGroup from "react-bootstrap/InputGroup";

class DatePicker extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      day: null,
      month: "month",
      year: null
    }
  }

  render() {
    return (
        <InputGroup>
          <FormControl placeholder="day" onChange={(event) => {
            this.setState({day: event.target.value})
          }}/>
          <select className="custom-select" onChange={(event) => {
            this.setState({month: event.target.value})
          }}>
            <option value="1">Jan</option>
            <option value="2">Feb</option>
            <option value="3">Mar</option>
            <option value="4">Apr</option>
            <option value="5">May</option>
            <option value="6">Jun</option>
            <option value="7">Jul</option>
            <option value="8">Aug</option>
            <option value="9">Sep</option>
            <option value="10">Oct</option>
            <option value="11">Nov</option>
            <option value="12">Dec</option>
          </select>
          <FormControl placeholder="year" onChange={(event) => {
            this.setState({year: event.target.value})
          }}/>
        </InputGroup>
    );
  }
}

export default DatePicker;
