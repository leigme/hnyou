{{define "header.tpl"}}
<meta charset="utf-8" />

<!-- Our CSS stylesheet files -->
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="/static/css/style.css">

<!-- Our Script files -->
<script src="http://cdn.staticfile.org/jquery/2.1.1-beta1/jquery.min.js"></script>  
<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<script type="text/javascript" src="/static/js/jquery.gallery.js"></script>
<script type="text/javascript" src="/static/js/modernizr.custom.53451.js"></script>
<script type="text/javascript">                
  $(function() {
    $('#banner').gallery({
      autoplay  : true
    });
  });
</script>
{{end}}