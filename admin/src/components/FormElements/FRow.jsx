import "./style.scss"


const FRow = ({ label, children }) => {
  return (
    <div className="FRow" >
      <div className="label">{label}:</div>
      <div className="component">{children}</div>
    </div>
  )
}

export default FRow
