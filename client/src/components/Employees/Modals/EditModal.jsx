import React, {useState} from "react";
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";

function EditModal(props) {
  const style = {
    color: "black"
  };

  const [name, setName] = useState(props.data["name"]);
  const [position, setPosition] = useState("");
  const [grade, setGrade] = useState("");
  const [tum, setTum] = useState(false);

  const data = {
    'id': props.data["id"],
    'name': name,
    'position': position,
    'grade': Number(grade),
    'tu-membership': tum,
  };

  return (
      <Modal style={style} show={props.show} onHide={props.hide}>
        <Modal.Header>
          <Modal.Title>Edit Employee</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group controlId="formName">
            <Form.Label>Name</Form.Label>
            <Form.Control placeholder="name" onChange={e => setName(e.target.value)} defaultValue={props.data["name"]}/>
          </Form.Group>
          <Form.Group controlId="formPosition">
            <Form.Label>Position</Form.Label>
            <Form.Control placeholder="position" onChange={e => setPosition(e.target.value)}
                          defaultValue={props.data["position"]}/>
          </Form.Group>
          <Form.Group controlId="formGrade">
            <Form.Label>Grade</Form.Label>
            <Form.Control placeholder="grade" onChange={e => setGrade(e.target.value)}
                          defaultValue={props.data["grade"]}/>
          </Form.Group>
          <Form.Group controlId="fromTum">
            <Form.Check type="checkbox" label="Trade Union Membership" onChange={() => setTum(!tum)}
                        defaultValue={props.data["tu-membership"]}/>
          </Form.Group>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="danger" onClick={() => props.handler(data)}>Create</Button>
          <Button variant="danger" onClick={props.hide}>Close</Button>
        </Modal.Footer>
      </Modal>
  );
}

export default EditModal;
