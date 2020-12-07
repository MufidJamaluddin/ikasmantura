import React, {PureComponent} from "react";

export default class Image extends PureComponent<any, {errored: boolean, src: string}>
{
    constructor(props)
    {
        super(props);

        this.state = {
            src: props.src,
            errored: false,
        };
    }

    onError = () =>
    {
        if (!this.state.errored)
        {
            this.setState({
                src: this.props.fallbackSrc,
                errored: true,
            });
        }
    }

    render()
    {
        const { src } = this.state;
        const {
            src: _1,
            fallbackSrc: _2,
            ...props
        } = this.props;

        return (
            <img
                src={src}
                onError={this.onError}
                alt={props.alt}
                {...props}
            />
        );
    }
}
