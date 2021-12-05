import "./style.scss"

const Button = ({children, className, ...props}) => {
  return (
    <button className={`Button ${className}`} {...props} >
      {children}
    </button>
  )
}

export default Button
