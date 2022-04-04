import graph from './js/graph.js';
import tree from './js/tree.js';

var hostPath = window.location.href;
var paths = hostPath.split("#");
// 这判断下，什么模式
if (0 < paths.length && "graph" == (paths[paths.length - 1]).toLowerCase()) {
    graph();
} else {
    tree();
}
