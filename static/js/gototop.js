window.onload = function(){
    var btn_gototop = document.getElementById('btn_gototop');
    var ph = document.documentElement.clientHeight;
    var timer = null;
    var isTop = true;

    window.onscroll = function(){
        var osTop = document.documentElement.scrollTop = document.body.scrollTop;
        if (osTop < ph) {
            btn_gototop.style.display = 'none';
        } else {
            btn_gototop.style.display = 'block';
        };
        if (!isTop) {
            clearInterval(timer);
        }
        isTop = false;
    }

    btn_gototop.onclick = function(){
        timer = setInterval(function(){
                var osTop = document.documentElement.scrollTop || document.body.scrollTop;
                var ispeed = Math.floor(-osTop / 5);
                document.documentElement.scrollTop = document.body.scrollTop = (osTop + ispeed);
                isTop = true;
            if (osTop == 0){
                clearInterval(timer)
            }
        },30);
    }    
}
