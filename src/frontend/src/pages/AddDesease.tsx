import React from 'react';
import { Button, Form } from 'react-bootstrap';
import MsgBox from '../components/MsgBox';

function AddDesease() {
    const [desease, setDesease] = React.useState("");
    const [sequence, setSequence] = React.useState("");
    const [show, setShow] = React.useState(false);
    const [title, setTitle] = React.useState("");
    const [text, setText] = React.useState("");

    const handleChangeFile = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files) {
            const file = event.target.files[0];
            const reader = new FileReader();
            reader.readAsText(file);
            reader.onload = () => {
                setSequence(reader.result as string);
            }
        }
    }

    const handleSubmitDisease = async (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        // Make post request to server
        try {
            const response = await fetch('/api/insert/', {
                method: 'POST',
                mode: 'same-origin',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    name: desease,
                    sequence: sequence
                })
            })
            
            if(response.ok){
                setTitle("Success");
                let result = await response.text()
                setText(result);
                setShow(true);
            } else{
                setTitle("Error");
                let result = await response.text()
                setText(result);
                setShow(true);
            }
        } catch (error) {
            setTitle('Error');
            setText('Internal Server error');
            setShow(true);
        }
    }

  return (
    <>
    <div className = 'container mt-5'>
        <div className='card'>
            <h1>Add Disease</h1>
            <Form className='form'>
                <p className='mt-2'>Disease</p>
                <input type="text" placeholder="Name" onChange={(e => setDesease(e.target.value)) } />
                <p>Sequence</p>
                <input type="file" onChange={handleChangeFile} />
                <Button variant="primary" onClick={handleSubmitDisease} className='button mt-2'>Submit</Button>
            </Form>
        </div>
    </div>
    {show && <MsgBox show={show} title={title} text={text} onHide={() => setShow(false)} />}
    </>
  )
}

export default AddDesease