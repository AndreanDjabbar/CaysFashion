export const AuthFormLayout = (props) => {
    return (
        <form className={props.className}>
            <p id="title">{props.titleForm}</p>
            {props.children}
            <div>
                <p id="guide">{props.guide} <a href={props.guideLink}>{props.guideType}</a> </p>
            </div>
        </form>
    )
}