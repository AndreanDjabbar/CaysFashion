export const Button = (props) => {
    return (
        <button 
        className={props.className}
        onClick={props.onClick}
        type={props.type}
        disabled={props.disabled}>
            {props.children}
        </button>
    )
}