import "./style.scss"

const Logo = ({size}) => {
  return (
    <div className={`Logo ${size}`} >
      <span>EASY</span>
      <span className="right-block" >CHEG<span className="orange-letter">%</span>RMA</span>
    </div>
  )
}

export default Logo
