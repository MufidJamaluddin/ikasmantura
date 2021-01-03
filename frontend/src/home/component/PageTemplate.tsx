import RegeTitle from "./RegeTitle";
import {Container} from "react-bootstrap";
import React, {createRef, PureComponent, RefObject} from "react";

class PageTitle extends PureComponent<any, {title, header, showTitle}> {

    constructor(props) {
        super(props);
        this.state = {
            title: '',
            header: '',
            showTitle: false,
        }
    }

    setHeader({ header, title, showTitle }) {
        this.setState({ header: header, title: title, showTitle: showTitle })
    }

    render() {
        let { title, header, showTitle } = this.state
        if(title) {
            document.title = title
        }
        return (
            <RegeTitle>
                { showTitle && title && (<h1 className="text-center display-4">{title}</h1>) }
                { header }
            </RegeTitle>
        )
    }
}

const ThemeContext = React.createContext({ setHeader: (props:any) => { } });

class PageTemplate extends PureComponent {

    constructor(props) {
        super(props);
        this.setHeader = this.setHeader.bind(this)
    }

    private titleRef: RefObject<PageTitle> = createRef()

    setHeader(headerParams: { header: any; title: any; showTitle: boolean; }) {
        if(this.titleRef.current) {
            this.titleRef.current.setHeader(headerParams)
        }
    }

    render() {
        return (
            <section className="features-icons bg-light">
                <PageTitle ref={this.titleRef} />
                <Container>
                    <ThemeContext.Provider value={{ setHeader: this.setHeader }}>
                        { this.props.children }
                    </ThemeContext.Provider>
                </Container>
            </section>
        )
    }
}

export default PageTemplate

export {
    ThemeContext
}
