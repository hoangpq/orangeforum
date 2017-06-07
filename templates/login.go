package templates

const loginsrc = `
{{ define "content" }}

<form action="/login" method="POST">
<input type="hidden" name="csrf" value="{{ .CSRF }}">
<input type="hidden" name="next" value="{{ .next }}">
Username: <input type="text" name="username" required>
Password: <input type="password" name="passwd" required>
<a href="/forgotpass">Forgot password?</a>
<input type="submit" value="Login">
</form>
{{ .Msg }}

{{ end }}`