export const OtpForm = (props) => {
    return (
        <div className={props.className}>
            <h1>{props.title}</h1>
            {props.children}
            <h3>{props.description}</h3>
        </div>
    )
}