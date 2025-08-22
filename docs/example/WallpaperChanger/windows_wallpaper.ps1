# This script downloads an image from a URL and sets it as the desktop wallpaper on Windows
# Set the below variables to customize the image URL, working directory, and wallpaper style
$IMAGE_URL = "http://example.com:8080/random"
$WORKDIR = "$env:userprofile\dailimage"
$WALLPAPER_STYLE = "Fit" # Fill, Fit, Stretch, Tile, Center, or Span

# Create WORKDIR if necessary
if ( -Not (Get-Item -Path "$WORKDIR") ) {
    New-Item -Path "$WORKDIR" -ItemType "Directory"
    Write-Output "Created '$WORKDIR'"
}

# Download image
$ProgressPreference = 'SilentlyContinue' # we set this because the progress bar makes downloads 100x slower
try {
    Invoke-WebRequest $IMAGE_URL -OutFile "$WORKDIR\background.png"
} catch {
    Write-Output "Error: failed to download $IMAGE_URL"
    exit 1
}
Write-Output "Downloaded '$IMAGE_URL' to '$WORKDIR\background.png'"
$ProgressPreference = 'Continue'

# Set background
# From Jose Espitia, https://www.joseespitia.com/2017/09/15/set-wallpaper-powershell-function/
Function Set-WallPaper {
 
<#
 
    .SYNOPSIS
    Applies a specified wallpaper to the current user's desktop
    
    .PARAMETER Image
    Provide the exact path to the image
 
    .PARAMETER Style
    Provide wallpaper style (Example: Fill, Fit, Stretch, Tile, Center, or Span)
  
    .EXAMPLE
    Set-WallPaper -Image "C:\Wallpaper\Default.jpg"
    Set-WallPaper -Image "C:\Wallpaper\Background.jpg" -Style Fit
  
#>
 
param (
    [parameter(Mandatory=$True)]
    # Provide path to image
    [string]$Image,
    # Provide wallpaper style that you would like applied
    [parameter(Mandatory=$False)]
    [ValidateSet('Fill', 'Fit', 'Stretch', 'Tile', 'Center', 'Span')]
    [string]$Style
)
 
$WallpaperStyle = Switch ($Style) {
  
    "Fill" {"10"}
    "Fit" {"6"}
    "Stretch" {"2"}
    "Tile" {"0"}
    "Center" {"0"}
    "Span" {"22"}
  
}
 
If($Style -eq "Tile") {
 
    New-ItemProperty -Path "HKCU:\Control Panel\Desktop" -Name WallpaperStyle -PropertyType String -Value $WallpaperStyle -Force
    New-ItemProperty -Path "HKCU:\Control Panel\Desktop" -Name TileWallpaper -PropertyType String -Value 1 -Force
 
}
Else {
 
    New-ItemProperty -Path "HKCU:\Control Panel\Desktop" -Name WallpaperStyle -PropertyType String -Value $WallpaperStyle -Force
    New-ItemProperty -Path "HKCU:\Control Panel\Desktop" -Name TileWallpaper -PropertyType String -Value 0 -Force
 
}
 
Add-Type -TypeDefinition @" 
using System; 
using System.Runtime.InteropServices;
  
public class Params
{ 
    [DllImport("User32.dll",CharSet=CharSet.Unicode)] 
    public static extern int SystemParametersInfo (Int32 uAction, 
                                                   Int32 uParam, 
                                                   String lpvParam, 
                                                   Int32 fuWinIni);
}
"@ 
  
    $SPI_SETDESKWALLPAPER = 0x0014
    $UpdateIniFile = 0x01
    $SendChangeEvent = 0x02
  
    $fWinIni = $UpdateIniFile -bor $SendChangeEvent
  
    $ret = [Params]::SystemParametersInfo($SPI_SETDESKWALLPAPER, 0, $Image, $fWinIni)
}
 
Set-WallPaper -Image "$WORKDIR\background.png" -Style $WALLPAPER_STYLE
Write-Output "Set wallpaper to '$WORKDIR\background.png' with style $WALLPAPER_STYLE"
