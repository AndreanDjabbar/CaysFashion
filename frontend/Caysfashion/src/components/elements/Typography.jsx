export const Title = (props) => {
    return (
        <p
        id={props.id}>{props.children}</p>
    )
}

export const Paragraph = (props) => {
    return (
        <p
        id={props.id}>{props.children}</p>
    )
}

export const Label = (props) => {
    return (
        <label
        id={props.id}
        htmlFor={props.htmlFor}>{props.children}</label>
    )
}