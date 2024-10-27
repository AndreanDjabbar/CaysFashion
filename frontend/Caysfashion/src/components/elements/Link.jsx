export const Link = (props) => {
    return (
        <a
        href={props.href}
        id={props.id}>{props.children}</a>
    )
}