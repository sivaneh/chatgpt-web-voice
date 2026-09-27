$root = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $root

Get-Content (Join-Path $root '.env') | ForEach-Object {
  $line = $_.Trim()
  if ($line -and -not $line.StartsWith('#') -and $line.Contains('=')) {
    $idx = $line.IndexOf('=')
    $name = $line.Substring(0, $idx).Trim()
    $value = $line.Substring($idx + 1).Trim()
    Set-Item -Path "Env:$name" -Value $value
  }
}

& (Join-Path $root 'bin\server.exe')
