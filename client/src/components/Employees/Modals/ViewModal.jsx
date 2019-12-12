import React from "react";
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";

function ViewModal(props) {
  const style = {
    color: "black"
  };

  return (
      <Modal style={style} show={props.show} onHide={props.hide}>
        <Modal.Header>
          <Modal.Title>View Employee</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <b>#{props.data["id"]}</b><br/>
          <b>Name:</b> {props.data["name"]}<br/>
          <b>Position:</b> {props.data["position"]}<br/>
          <b>Grade:</b> {props.data["grade"]}<br/>
          <b>{props.data["tu-membership"] ? "Is" : "Is not"} Trade Union member</b><br/>
          {/*TODO: Salary*/}
        </Modal.Body>
        <Modal.Footer>
          <Button variant="danger" onClick={props.hide}>Close</Button>
        </Modal.Footer>
      </Modal>
  );
}

export default ViewModal;
