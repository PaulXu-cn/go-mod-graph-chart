import * as d3 from 'd3';
import CSS from '../css/styles.css';

var treeWidth = 1;
var treeDepth = 1;

var width = 4000;
var height = 2500;

var innerWidth = 0
var innerHeight = 0;

var subWidth = 640;
var subHeight = 480;

var root;
var color;
var anTree;

const margin = { top: 5, right: 200, bottom: 5, left: 200 };

export default function tree() {
    const svg = d3.select('#app').append('svg').attr('id', 'svg')
        .attr('width', width).attr('height', height);
    const subSvg = d3.select('#sub-board').append('svg').attr('id', 'sub-svg')
        .attr('width', subWidth).attr('height', subHeight);
    
    let markerBoxWidth = 10;
    let markerBoxHeight = 10;
    let refX = 10;
    let refY = 4;
        svg.append('defs')
        .append('marker')
        .attr('id', 'arrow')
        .attr('viewBox', [0, 0, markerBoxWidth, markerBoxHeight])
        .attr('refX', refX)
        .attr('refY', refY)
        .attr('markerWidth', markerBoxWidth)
        .attr('markerHeight', markerBoxHeight)
        .attr('orient', 'auto-start-reverse')
        .append('path')
        .attr('d', d3.line()([[0, 0], [8, 4], [0, 8], [0, 7], [9, 4], [0, 1], [0, 0]]))
        .attr('stroke', 'black');
    

    const g = svg.append('g')
        .attr('transform', `translate(${margin.left}, ${margin.top})`);

    const fill = d => {
        if (d.depth === 0)
            return color(d.data.name)
        else
            return color(d.parent.data.name)
    }

    document.getElementById('mask').addEventListener('click', function (ev) {
        if ('mask' == ev.target.getAttribute('id')) {
            // ev.target.setAttribute("class", "mask hide");
            document.getElementById('mask').setAttribute("class", "mask hide");
        }
    });

    const subRender = function (data) {
        let anotation = 500;
        color = d3.scaleOrdinal()
            .domain(data.descendants().filter(d => d.depth <= 1).map(d => d.data.name))
            .range(d3.schemeCategory10);

        subSvg.selectAll("path")
            .data(data.links())
            .join("path")
            .classed('tree-path', true)
            .attr("d", d3.linkHorizontal().x(d => anotation - d.y).y(d => d.x));

        subSvg.selectAll('circle').data(data.descendants()).join('circle')
            // optionally, we can use stroke-linejoin to beautify the path connection; 
            //.attr("stroke-linejoin", "round")
            .classed('tree-circle', true)
            .attr("fill", fill)
            .attr('id', (d, i) => `circle-${i}`)
            .attr('cx', d => anotation - d.y)
            .attr('cy', d => d.x)
            .attr("r", function (d) {
                if (d.children) {
                    return 6 + Math.sqrt(d.children.length) * 3;
                } else {
                    return 7;
                }
            });

        subSvg.selectAll('text').data(data.descendants()).join('text')
            .classed('tree-text', true)
            .attr("text-anchor", d => d.children ? "end" : "start")
            .attr('id', (d, i) => `text-${i}`)
            // note that if d is a child, d.children is undefined which is actually false! 
            .attr('x', d => (d.children ? -6 : 6) + (anotation - d.y))
            .attr('y', function (d) { return d.depth % 2 ? d.x - 3 : d.x + 14 })
            .text(function (d) {
                // 默认最多只展示版本号，不展示hash
                return d.data.name.replace(/(@[\w\.]*?)(-)(.*$)/, "$1")
            });
    }

    const renderSubTree = function (d, i) {
        let theName = d.data.name;
        let subTree = undefined;
        if (anTree[theName] !== undefined) {
            subTree = anTree[theName]
        } else {
            return
        }

        document.getElementById('mask').setAttribute("class", "mask");

        let theSubHeight = subTree.children.length * 150;
        d3.select('#sub-board').select('#sub-svg').attr('width', subWidth).attr('height', theSubHeight);

        let subRoot = d3.hierarchy(subTree);
        subRoot = d3.tree().size([theSubHeight, subWidth - 200])(subRoot);
        subRender(subRoot);
    }

    const render = function (data, anData) {
        color = d3.scaleOrdinal()
            .domain(data.descendants().filter(d => d.depth <= 1).map(d => d.data.name))
            .range(d3.schemeCategory10);

        g.selectAll("path")
            .data(data.links())
            .join("path")
            .classed('tree-path', true)
            .attr("d", d3.linkHorizontal().x(d => d.y).y(d => d.x))
            .attr("marker-end", function (d, i) {
                let check = anData[d.target.data.name];
                if (undefined != check) {
                    return "url(#arrow)"
                }
            });

        g.selectAll('circle').data(data.descendants()).join('circle')
            // optionally, we can use stroke-linejoin to beautify the path connection; 
            //.attr("stroke-linejoin", "round")
            .classed('tree-circle', true)
            .attr("fill", fill)
            .attr('id', (d, i) => `circle-${i}`)
            .attr('cx', d => d.y)
            .attr('cy', d => d.x)
            .attr("r", function (d) {
                if (d.children) {
                    return 6 + Math.sqrt(d.children.length) * 3;
                } else {
                    return 7;
                }
            }).on('mouseover', function (d, i) {
                d3.select(this).classed('on', true);
                d3.select("#text-" + i).classed('on', true);
            }).on('mouseout', function (d, i) {
                d3.select(this).classed('on', false).attr('stroke', fill(d));
                d3.select("#text-" + i).classed('on', false);
            }).on('click', renderSubTree);

        g.selectAll('text').data(data.descendants()).join('text')
            .classed('tree-text', true)
            .attr("text-anchor", d => d.children ? "end" : "start")
            .attr('id', (d, i) => `text-${i}`)
            // note that if d is a child, d.children is undefined which is actually false! 
            .attr('x', d => (d.children ? -6 : 6) + d.y)
            .attr('y', function (d) { return d.depth % 2 ? d.x - 3 : d.x + 14 })
            .text(function (d) {
                // 默认最多只展示版本号，不展示hash
                return d.data.name.replace(/(@[\w\.]*?)(-)(.*$)/, "$1")
            })
            .on('mouseover', function (d, i) {
                // 鼠标放上去时，才展示后面的hash值
                d3.select(this).classed('on', true).text(d => d.data.name);
                d3.select("#circle-" + i).classed('on', true);
            }).on('mouseout', function (d, i) {
                d3.select(this).classed('on', false).text(function (d) {
                    return d.data.name.replace(/(@[\w\.]*?)(-)(.*$)/, "$1")
                });
                d3.select("#circle-" + i).classed('on', false).attr('stroke', fill(d));
            }).on('click', renderSubTree);
    }

    d3.json('/an-tree.json').then(data => {
        anTree = data.data.tree;

        d3.json('/tree.json').then(data => {
            let treeData = data.data.tree;
            treeWidth = data.data.width;
            treeDepth = data.data.depth;

            innerWidth = treeDepth * 450;
            innerHeight = treeWidth * 70;

            width = innerWidth + margin.left + margin.right;
            height = innerHeight + margin.top + margin.bottom;

            d3.select('#app').select('svg').attr('width', width).attr('height', height);

            root = d3.hierarchy(treeData);
            // alternatively, we can set size of each node; 
            // root = d3.tree().nodeSize([30, width / (root.height + 1)])(root);
            root = d3.tree().size([innerHeight, innerWidth])(root);
            render(root, anTree);

            setTimeout(function () {
                window.scrollTo(0, height / 2)
            }, 1000);
        });
    });

}
