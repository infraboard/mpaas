// 深色主题
var Solarized_Darcula = {
    foreground: '#d2d8d9',
    background: '#3d3f41',
    cursor: '#d0d0d0',

    black: '#25292a',
    brightBlack: '#25292a',

    red: '#f24840',
    brightRed: '#f24840',

    green: '#629655',
    brightGreen: '#629655',

    yellow: '#b68800',
    brightYellow: '#b68800',

    blue: '#2075c7',
    brightBlue: '#2075c7',

    magenta: '#797fd4',
    brightMagenta: '#797fd4',

    cyan: '#15968d',
    brightCyan: '#15968d',

    white: '#d2d8d9',
    brightWhite: '#d2d8d9'
}

// 浅色主题
var GitHub = {
    foreground: '#3e3e3e',
    background: '#f4f4f4',
    cursor: '#3f3f3f',

    black: '#3e3e3e',
    brightBlack: '#666666',

    red: '#970b16',
    brightRed: '#de0000',

    green: '#07962a',
    brightGreen: '#87d5a2',

    yellow: '#f8eec7',
    brightYellow: '#f1d007',

    blue: '#003e8a',
    brightBlue: '#2e6cba',

    magenta: '#e94691',
    brightMagenta: '#ffa29f',

    cyan: '#89d1ec',
    brightCyan: '#1cfafe',

    white: '#ffffff',
    brightWhite: '#ffffff'
}

var getTermSize = (term) => {
    const MINIMUM_COLS = 2;
    const MINIMUM_ROWS = 1;
    const core = term._core;
    const dims = core._renderService.dimensions;

    if (dims.css.cell.width === 0 || dims.css.cell.height === 0) {
        return undefined;
    }

    const scrollbarWidth = term.options.scrollback === 0 ? 0 : core.viewport.scrollBarWidth;
    const parentElementStyle = window.getComputedStyle(term.element.parentElement);
    const parentElementHeight = parseInt(parentElementStyle.getPropertyValue('height'));
    const parentElementWidth = Math.max(0, parseInt(parentElementStyle.getPropertyValue('width')));
    const elementStyle = window.getComputedStyle(term.element);
    const elementPadding = {
        top: parseInt(elementStyle.getPropertyValue('padding-top')),
        bottom: parseInt(elementStyle.getPropertyValue('padding-bottom')),
        right: parseInt(elementStyle.getPropertyValue('padding-right')),
        left: parseInt(elementStyle.getPropertyValue('padding-left'))
    };
    const elementPaddingVer = elementPadding.top + elementPadding.bottom;
    const elementPaddingHor = elementPadding.right + elementPadding.left;
    const availableHeight = parentElementHeight - elementPaddingVer;
    const availableWidth = parentElementWidth - elementPaddingHor - scrollbarWidth;
    const geometry = {
        cols: Math.max(MINIMUM_COLS, Math.floor(availableWidth / dims.css.cell.width)),
        rows: Math.max(MINIMUM_ROWS, Math.floor(availableHeight / dims.css.cell.height))
    };
    return geometry
}