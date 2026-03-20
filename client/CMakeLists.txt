cmake_minimum_required( VERSION 3.15 )
project( Phantom-Operator )

##
## cmake sets
##
set( QT_VERSION 5 )
set( CMAKE_CXX_STANDARD 20 )
set( CMAKE_AUTOMOC ON )
set( CMAKE_AUTORCC ON )
set( CMAKE_AUTOUIC ON )
set( CMAKE_CPP_COMPILER /usr/bin/x86_64-w64-mingw32-g++ )
set( CMAKE_C_COMPILER /usr/bin/x86_64-w64-mingw32-gcc )
set( CMAKE_CXX_FLAGS "-fpermissive" )
set( REQUIRED_LIBS Core Gui Widgets Network WebSockets Sql )
set( REQUIRED_LIBS_QUALIFIED Qt5::Core Qt5::Gui Qt5::Widgets Qt5::Network Qt5::WebSockets Qt5::Sql )
set( APP_ICON_RESOURCE_WINDOWS "data/Phantom-Operator.qrc" )
set( CMAKE_RUNTIME_OUTPUT_DIRECTORY .. )
set( CMAKE_INCLUDE_CURRENT_DIR ON )

##
## Ensure 3.10 is used if present.
## If not, modules will not work within the client.
##
set( Python_ADDITIONAL_VERSIONS 3.10 )

##
## import includes and external libraries
##
include_directories( include  )
include_directories( external/spdlog/include )
include_directories( external/json/include )
include_directories( external/toml )

##
## find qt packages
##
find_package( Qt${QT_VERSION} COMPONENTS ${REQUIRED_LIBS} REQUIRED )

##
## apple support
##
if(APPLE)
  find_package( Python 3 COMPONENTS Interpreter Development REQUIRED )
  set( PYTHON_MAJOR $ENV{Python_VERSION_MAJOR} )
  set( PYTHON_MINOR $ENV{Python_VERSION_MINOR} )
  set( PYTHONLIBS_VERSION_STRING ${Python_VERSION} )
  set( PYTHON_INCLUDE_DIR ${Python_INCLUDE_DIRS} )
  set( PYTHON_LIBRARIES ${Python_LIBRARIES} )
  message( "Apple - Using Python:${Python_VERSION_MAJOR} - Libraries:${PYTHON_LIBRARIES} - includeDirs: ${PYTHON_INCLUDE_DIR}" )
elseif(UNIX)
  find_package( PythonLibs 3 REQUIRED )
else()
  set( PYTHONLIBS_VERSION_STRING $ENV{PY_VERSION} )
endif()



if(APPLE)
    execute_process( COMMAND brew --prefix OUTPUT_VARIABLE BREW_PREFIX ) #this because brew install location differs Intel/Apple Silicon macs
    string( STRIP ${BREW_PREFIX} BREW_PREFIX ) #for some reason this happens: https://gitlab.kitware.com/cmake/cmake/-/issues/22404
    include_directories( "${BREW_PREFIX}/bin/python3.10" )
    include_directories( "${BREW_PREFIX}/Frameworks/Python.framework/Headers" )
elseif(UNIX)
    include_directories( ${PYTHON_INCLUDE_DIRS} )
endif()

set( HAVOC_INCLUDE

        # Misc
        include/Phantom-Operator/Connector.hpp
        include/Phantom-Operator/DBManager/DBManager.hpp
        include/Phantom-Operator/DemonCmdDispatch.h
        include/Phantom-Operator/Packager.hpp
        include/Phantom-Operator/Phantom-Operator.hpp
        include/Phantom-Operator/Service.hpp
        include/Phantom-Operator/CmdLine.hpp

        # Python Api
        include/Phantom-Operator/PythonApi/PythonApi.h
        include/Phantom-Operator/PythonApi/Event.h
        include/Phantom-Operator/PythonApi/Phantom-OperatorUi.h
        include/Phantom-Operator/PythonApi/PyAgentClass.hpp
        include/Phantom-Operator/PythonApi/UI/PyDialogClass.hpp
        include/Phantom-Operator/PythonApi/UI/PyLoggerClass.hpp
        include/Phantom-Operator/PythonApi/UI/PyTreeClass.hpp
        include/Phantom-Operator/PythonApi/UI/PyWidgetClass.hpp

        # Dialogs
        include/UserInterface/Dialogs/Payload.hpp
        include/UserInterface/Dialogs/About.hpp
        include/UserInterface/Dialogs/Connect.hpp
        include/UserInterface/Dialogs/Listener.hpp

        # small widgets
        include/UserInterface/SmallWidgets/EventViewer.hpp
        include/UserInterface/SmallWidgets/EventViewer.hpp

        # widgets
        include/UserInterface/Widgets/Chat.hpp
        include/UserInterface/Widgets/DemonInteracted.h
        include/UserInterface/Widgets/ListenerTable.hpp
        include/UserInterface/Widgets/ProcessList.hpp
        include/UserInterface/Widgets/PythonScript.hpp
        include/UserInterface/Widgets/SessionTable.hpp
        include/UserInterface/Widgets/SessionGraph.hpp
        include/UserInterface/Widgets/TeamserverTabSession.h
        include/UserInterface/Widgets/ScriptManager.h
        include/UserInterface/Widgets/LootWidget.h
        include/UserInterface/Widgets/FileBrowser.hpp
        include/UserInterface/Widgets/Teamserver.hpp
        include/UserInterface/Widgets/Store.hpp

        include/UserInterface/Phantom-OperatorUI.hpp
)

# src for UI
set(
        HAVOC_UI

        src/UserInterface/Phantom-OperatorUi.cc
        src/global.cc

        # Dialogs
        src/UserInterface/Dialogs/About.cc
        src/UserInterface/Dialogs/Connect.cc
        src/UserInterface/Dialogs/Listener.cc
        src/UserInterface/Dialogs/Payload.cc

        # Widgets
        src/UserInterface/Widgets/SessionTable.cc
        src/UserInterface/Widgets/SessionGraph.cc
        src/UserInterface/Widgets/Chat.cc
        src/UserInterface/Widgets/ListenersTable.cc
        src/UserInterface/Widgets/DemonInteracted.cc
        src/UserInterface/Widgets/TeamserverTabSession.cc
        src/UserInterface/Widgets/PythonScript.cc
        src/UserInterface/Widgets/ScriptManager.cc
        src/UserInterface/Widgets/LootWidget.cc
        src/UserInterface/Widgets/FileBrowser.cc
        src/UserInterface/Widgets/Teamserver.cc
        src/UserInterface/Widgets/Store.cc
        src/UserInterface/Widgets/ProcessList.cc

        # SmallWidgets
        src/UserInterface/SmallWidgets/EventViewer.cc
)

set(
        HAVOC_UTIL

        src/Util/ColorText.cpp
        src/Util/Base64.cpp
        src/Util/Base.cpp
)

# src for Phantom-Operator
set( HAVOC_SOURCE
        src/Main.cc

        src/Phantom-Operator/Packager.cc
        src/Phantom-Operator/Connector.cc
        src/Phantom-Operator/Service.cc

        src/Phantom-Operator/DBManger/DBManager.cc
        src/Phantom-Operator/DBManger/Teamserver.cc
        src/Phantom-Operator/DBManger/Scripts.cc

        src/Phantom-Operator/Demon/ConsoleInput.cc
        src/Phantom-Operator/Demon/CommandSend.cc
        src/Phantom-Operator/Demon/CommandOutput.cc
        src/Phantom-Operator/Demon/Commands.cc

        src/Phantom-Operator/PythonApi/PythonApi.cc
        src/Phantom-Operator/PythonApi/Phantom-Operator.cc
        src/Phantom-Operator/PythonApi/Phantom-OperatorUi.cc
        src/Phantom-Operator/PythonApi/UI/PyDialogClass.cc
        src/Phantom-Operator/PythonApi/UI/PyLoggerClass.cc
        src/Phantom-Operator/PythonApi/UI/PyTreeClass.cc
        src/Phantom-Operator/PythonApi/UI/PyWidgetClass.cc
        src/Phantom-Operator/PythonApi/PyDemonClass.cc
        src/Phantom-Operator/PythonApi/Event.cc
        src/Phantom-Operator/PythonApi/PyAgentClass.cc

        src/Phantom-Operator/Phantom-Operator.cc

    ${HAVOC_UI}
)

add_executable( ${PROJECT_NAME} ${HAVOC_INCLUDE} ${HAVOC_SOURCE} ${HAVOC_UTIL} ${APP_ICON_RESOURCE_WINDOWS} data/Phantom-Operator.rc )
add_definitions( -DQT_NO_DEBUG_OUTPUT )

target_link_libraries(
    ${PROJECT_NAME}
    ${REQUIRED_LIBS_QUALIFIED}
    ${PYTHON_LIBRARIES}
    # spdlog::spdlog
    # spdlog::spdlog_header_only
)
