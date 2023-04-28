document.addEventListener('astilectron-ready', function() {
    // Send
    document.getElementById('test_button').onclick = function() {
        astilectron.sendMessage("event.test_button.clicked", null);
    };

    // Recive
    astilectron.onMessage(function(msg) {
        switch(msg) {
        case 'event.golang.onwait':
            return '{"msg":"OK"}';
        default:
            break;
        };
    });
});