#!/bin/bash
##########################################################################################
# build.sh
##########################################################################################

SCRIPT_FILENAME_THIS=${0##*/}

usage ()
{
	echo ''
	echo 'Cross compiling script of Go server'
	echo ''
	echo 'Usage : '${SCRIPT_FILENAME_THIS}' [-h] [-os <OS> [-arch <Architecture>]]'
	echo ''
	echo '<Optional switches>'
	echo '  -h/--help   Show help'
	echo '  -os         Set cross compiling target os.'
	echo '              Available OS type:              darwin, linux, windows'
	echo '  -arch       Set cross compiling target architecture.'
	echo '              Currently arm architecture are supported in linux only!'
	echo '              Avialable Architecture:         386, amd64, arm5, arm6, arm7, arm64'
	echo ''
	echo '<Example>'
	echo '  Cross compile for Linux x64 target:'
	echo '              '${SCRIPT_FILENAME_THIS}' -os linux -arch amd64'
	echo '  Local compile for Linux with default architecture:'
	echo '              '${SCRIPT_FILENAME_THIS}' -os linux'
	echo '  MacOSX with intel 386 architecture:'
	echo '              '${SCRIPT_FILENAME_THIS}' -os darwin -arch 386'
	echo '  Local compile for local os target, no switch needed:'
	echo '              '${SCRIPT_FILENAME_THIS}
	echo ''

	exit
}

# Building directory settings
CURR_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Switch to the building directory
cd "$CURR_DIR"

# Check parameters
#if [ "$1" = "" ];
#then
#	usage
#fi

PARAM_OS=""
PARAM_ARCH=""
while [ "$1" != "" ]; do
	case $1 in
        "-os")         shift
                       PARAM_OS=$1
                       ;;
        "-arch")       shift
                       PARAM_ARCH=$1
                       ;;
        "-h")          usage
                       ;;
        "--help")      usage
                       ;;
	esac
	shift
done

# Set GOOS env variable
case $PARAM_OS in
	darwin )       ;;
	linux )        ;;
	windows )      ;;
	"" )           ;;
	* )            echo "Invalid -os parameters!"
	               usage
	               ;;
esac
if !([ "$PARAM_OS" = "" ]);
then
	echo "set GOOS variable to \"$PARAM_OS\""
    export GOOS="$PARAM_OS"
fi
# Set GOARCH env variable
if !([ "$PARAM_OS" = "" ]);
then
	case $PARAM_OS in
		darwin )
			case $PARAM_ARCH in
				386 )		;;
				amd64 )		;;
				* )       	echo "Invalid -arch parameters for darwin os.Only 386 and amd64 are supported!"
							usage
							;;
			esac
			;;
		linux )
			case $PARAM_ARCH in
				386 )		;;
				amd64 )		;;
				arm5 )		;;
				arm6 )		;;
				arm7 )		;;
				arm64 )		;;
				* )       	echo "Invalid -arch parameters for linux os!"
							usage
							;;
			esac
			;;
		windows )
			case $PARAM_ARCH in
				386 )		;;
				amd64 )		;;
				* )       	echo "Invalid -arch parameters for windows os.Only 386 and amd64 are supported!"
							usage
							;;
			esac
			;;
	esac
	if !([ "$PARAM_ARCH" = "" ]);
	then
		echo "set GOARCH variable to \"$PARAM_ARCH\""
		export GOARCH="$PARAM_ARCH"
	fi
fi



go build -x

rm -rf ./bin
mkdir -p bin
mv -f ./testwywlogin bin/



