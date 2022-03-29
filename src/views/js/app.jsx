var App = React.createClass({
    getInitialState: function() {
        return {
            list: [],
            content: ""
        };
    },


    handleClick: function(e) {

        var bodyFormData = new FormData();
        var self = this;

        bodyFormData.append('book_text', this.state.content);

        axios({
            method: "post",
            url: "http://localhost:8080/book_text",
            data: bodyFormData,
            headers: { "Content-Type": "application/json" },
          })
            .then(function (response) {
              //handle success
              console.log(response.data);
              self.setState({ list: response.data })
            })
            .catch(function (response) {
              //handle error
              console.log(response);
            });
    },

    handleChange: function(event) {
        this.setState({ content: event.target.value});
        console.log(this.state.content);
    },

    render: function() {
        return (<div className="container">
            <div className="row" style={{ display: "flex"}}>
                <div className="col-md-6">
                    <textarea class="form-control" placeholder="Place your text here!" value={this.state.content} onChange={(e) => this.handleChange(e)} multiple rows="22" cols="70"/>
                </div>
                <div className="col-md-3">
                    <ul className="list-group">
                        {this.state.list.map(word => 
                            (<li className="list-group-item" key={word.word}>
                                <div className="row" >
                                    <div className="col-md-8">
                                        {word.word}
                                    </div>
                                    <div className="col-md-4">
                                        {word.count}
                                    </div>
                                </div>
                            </li>)
                        )}
                    </ul>
                </div>
            </div>
            <div style={{ "text-align": "center"}}>
                <button type="button" className="btn btn-primary" onClick={(e) => this.handleClick(e)}>PARSE</button>
            </div>
        </div>);
    }
});

React.render(
    React.createElement(App, {}),
    document.body
);
