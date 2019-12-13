import React from "react";
import FormControl from "react-bootstrap/FormControl";
import InputGroup from "react-bootstrap/InputGroup";

function DatePicker(props) {
  return (
      <InputGroup>
        <FormControl placeholder="day" onChange={e => props.setters.setDay(e.target.value)} defaultValue={1}/>
        <select className="custom-select" onChange={e => props.setters.setMonth(e.target.value)}>
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
        <FormControl placeholder="year" onChange={e => props.setters.setYear(e.target.value)} defaultValue={1}/>
      </InputGroup>
    );
}

export default DatePicker;
