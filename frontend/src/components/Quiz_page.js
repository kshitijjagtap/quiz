import { useState } from 'react'
import { quiz } from './quiz'
import ResultPage from './ResultPage'

export default function Quiz_page() {
  const [activeQuestion, setActiveQuestion] = useState(0)
  const [answer,setAnswer]=useState('')
  const [indexofanswer,setIndexofanswer]=useState(0)
 const [marks,setMarks]=useState(null)
 const [resultpage,setResultpage]=useState(false)
  const { questions } = quiz
  const { question} = questions[activeQuestion]
 

  const onClickNext = () => {              /////// on clicking on next
    if (activeQuestion !== questions.length - 1) {
      setActiveQuestion((prev) => prev + 1)
    } 
   

  

   
      
     if(indexofanswer<=questions.length){        /////////////////011111111111111
      let ind = indexofanswer.toString()
      const data ={
        Id: ind,
        Ans: answer,
     
       }

      fetch(' http://localhost:9000/api/answer', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({data})
      })
    }

     if(indexofanswer===questions.length-1){ ////////////2222222222222222222222
      
      const data ={
        username:"kshitijjagtap",
        marks:"",
        time:"",
      }


     fetch(' http://localhost:9000/api/submit', {
       method: 'POST',
       headers: {
         'Content-Type': 'application/json'
       },
       body: JSON.stringify({data})
     })
     .then((response) => response.json())
     .then((data) => {
       const marks = data['You got total marks:']
       setMarks(marks)
     })
     .catch((error) => console.error(error))
    
     setIndexofanswer((prev) => prev + 9)
     setResultpage(true)
  
    }
    setIndexofanswer((prev) => prev + 1)
     setAnswer('')
  } 

  const handleanswerchange = (event)=>{
    setAnswer(event.target.value);
  }

 

 

  return (
    <>
   { resultpage ?(<ResultPage marks={marks}/>):
         <div className='main-div'>
          <div className='sub-div'>
          <h2>QUIZ APP</h2>
          <h2>{question}</h2>
         
           <input className='user-input'  value={answer} onChange={handleanswerchange} placeholder="Answer"/>
          
          
            <button className='btn'
              onClick={onClickNext}
           >
              {activeQuestion === questions.length - 1 ? 'Finish' : 'Next'}
              
            </button>
            </div>
        </div>
       } 
        </>
      )
 
}
