import React from 'react'
import { Button, Form } from 'react-bootstrap'
import MsgBox from '../components/MsgBox'

function DNATest() {
    const [name, setName] = React.useState('')
    const [sequence, setSequence] = React.useState('')
    const [disease, setDisease] = React.useState('')
    const [selected, setSelected] = React.useState('bm')
    const [show, setShow] = React.useState(false)
    const [title, setTitle] = React.useState('')
    const [text, setText] = React.useState('')

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

    const handleOptionChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSelected(event.target.value)
    }

    const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        try {
            const response = await fetch('/api/match/'+selected, {
                method: 'POST',
                mode: 'same-origin',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "name": name,
                    "sequence": sequence,
                    "penyakit": disease
                })
            })
            
            if(response.ok){
                setTitle("Success");
                let result = await response.text()
                let arr = result.split(' ')
                let temp = "Date: " + arr[0] + " " + arr[1] + " " + arr[2] + "\n" +
                            "Name: " + arr[3] + "\n" +
                            "Disease: " + arr[4] + "\n" +
                            "Verdict: " + arr[6] + "\n" +
                            "Similarity: " + arr[5]
                setText(temp);
                setShow(true);
            } else{
                setTitle("Error");
                let result = await response.text()
                setText(result);
                setShow(true);
            }
        } catch (error) {
            setTitle("Error");
            setText("Internal Server error");
            setShow(true);

        }
    }

  return (
    <>
    <div className = 'container mt-5'>
        <div className='card'>
            <h1>DNA Test</h1>
            <Form className='form'>
                <p className='mt-2'>Username</p>
                <input type="text" placeholder="Name" onChange={e => setName(e.target.value)}/>
                <p>Sequence</p>
                <input type="file" onChange={handleChangeFile}/>
                <p>Disease</p>
                <input type="text" placeholder='Disease' onChange={e => setDisease(e.target.value)}/>
                <p>Algorithm</p>
                <div className='input-radio'>
                    <label>
                    <input type="radio" name="algorithm" value="bm" checked={selected === 'bm'} onChange={handleOptionChange}/>
                    Boyer-Moore
                    </label>
                    <label>                 
                    <input type="radio" name="algorithm" value="kmp" checked={selected === 'kmp'} onChange={handleOptionChange}/>
                    Knuth-Morris-Pratt
                    </label>
                </div>
                <Button variant="primary" onClick={handleSubmit} className='button mt-2'>Submit</Button>
            </Form>
        </div>
    </div>
    {show && <MsgBox show={show} title={title} text={text} onHide={() => setShow(false)} />}
    </>
  )
}

export default DNATest